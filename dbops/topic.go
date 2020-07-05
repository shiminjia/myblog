package dbops

import (
	"database/sql"
	"log"
)

type Topic struct {
	Id           int
	Title        string
	Label        string
	CategoryId   int
	CategoryName string
	Content      string
	Status       int
	Time         *Time
}

func (topic *Topic) GetTopicAmount(categoryIdInt int) (int64, error) {
	var newSql string
	if categoryIdInt == 0 {
		newSql = "select count(*) from topics where status <> 9"
	} else {
		newSql = "select count(*) from topics where category_id = ? and status <> 9"
	}
	stmtOut, err := DB.Prepare(newSql)
	if err != nil {
		return -1, err
	}
	defer stmtOut.Close()
	var TopicAmount int64
	if categoryIdInt == 0 {
		err = stmtOut.QueryRow().Scan(&TopicAmount)
	} else {
		err = stmtOut.QueryRow(categoryIdInt).Scan(&TopicAmount)
	}
	if err != nil {
		return -1, err
	}
	return TopicAmount, nil
}

func (topic *Topic) CreateTopic() (int64, error) {
	stmtOut, err := DB.Prepare("insert into topics (title, label, category_id, content, status, created_at, updated_at) value (?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmtOut.Close()

	topic.Status = 1
	time := NewTime()
	result, err := stmtOut.Exec(topic.Title, topic.Label, topic.CategoryId, topic.Content, topic.Status, time.CreatedAt, time.UpdatedAt)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (topic *Topic) GetTopicById() (*Topic, error) {
	err := DB.QueryRow("select title, label, category_id, content, status from topics "+
		"where id = ? and status <> 9", topic.Id).Scan(&topic.Title, &topic.Label, &topic.CategoryId, &topic.Content, &topic.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return topic, nil
}

//func (topic *Topic) GetOneTopicById() (*Topic, error) {
//	err := DB.QueryRow("select title, label, category_id, content, status, created_at, updated_at "+
//		"from topics where id = ? and status <> 9", topic.Id).Scan(&topic.Title, &topic.Label, &topic.CategoryId,
//		&topic.Content, &topic.Status, &topic.Time.CreatedAt, &topic.Time.UpdatedAt)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	return topic, nil
//}

func (topic *Topic) GetOneTopicById() (*Topic, error) {
	err := DB.QueryRow(`select tp.title, tp.label, cg.name, tp.content, tp.status, tp.created_at, tp.updated_at
		from topics tp
		inner join categories cg
		on tp.category_id = cg.id
		where tp.id = ? and tp.status <> 9`, topic.Id).Scan(&topic.Title, &topic.Label, &topic.CategoryName,
		&topic.Content, &topic.Status, &topic.Time.CreatedAt, &topic.Time.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return topic, nil
}

func (topic *Topic) GetSomeTopicsByCategoryId(categoryId int, page int, limit int) ([]*Topic, error) {
	newSql := `select tp.id, tp.title, tp.label, tp.category_id, cg.name, tp.content, tp.status, tp.created_at, tp.updated_at
		from topics tp
		inner join categories cg
		on tp.category_id = cg.id
		where tp.status <> 9 `
	if categoryId != 0 {
		newSql += `and cg.id = ? `
	}
	newSql += `order by tp.id desc limit ?, ?`

	stmtOut, err := DB.Prepare(newSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmtOut.Close()

	var rows *sql.Rows
	if categoryId == 0 {
		rows, err = stmtOut.Query((page-1)*limit, limit)
	} else {
		rows, err = stmtOut.Query(categoryId, (page-1)*limit, limit)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []*Topic
	for rows.Next() {
		topic := Topic{
			Time: &Time{},
		}
		rows.Scan(&topic.Id, &topic.Title, &topic.Label, &topic.CategoryId, &topic.CategoryName, &topic.Content,
			&topic.Status, &topic.Time.CreatedAt, &topic.Time.UpdatedAt)
		result = append(result, &topic)
	}

	return result, nil
}

//func (topic *Topic) GetSomeTopics(page int, limit int) ([]*Topic, error) {
//	stmtOut, err := DB.Prepare(`
//        select tp.id, tp.title, tp.label, tp.category_id, cg.name, tp.content, tp.status, tp.created_at, tp.updated_at
//        from topics tp
//        inner join categories cg
//        on tp.category_id = cg.id
//		where tp.status <> 9
//		order by tp.id desc
//		limit ?, ?`)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	defer stmtOut.Close()
//
//	rows, err := stmtOut.Query((page-1)*limit+1, page*limit)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	var result []*Topic
//	for rows.Next() {
//		topic := Topic{
//			Time: &Time{},
//		}
//		rows.Scan(&topic.Id, &topic.Title, &topic.Label, &topic.CategoryId, &topic.CategoryName, &topic.Content,
//			&topic.Status, &topic.Time.CreatedAt, &topic.Time.UpdatedAt)
//		result = append(result, &topic)
//	}
//
//	return result, nil
//}

func (topic *Topic) UpdateTopicById() (int64, error) {
	stmtOut, err := DB.Prepare("update topics set title = ?, label = ?, category_id = ?, content = ?, updated_at = ? where id = ?")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	time := NewTime()
	result, err := stmtOut.Exec(topic.Title, topic.Label, topic.CategoryId, topic.Content, time.UpdatedAt, topic.Id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (topic *Topic) DeleteTopicById() (int64, error) {
	stmtOut, err := DB.Prepare("update topics set status = ?, deleted_at = ? where id = ? and status <> 9")
	if err != nil {
		log.Println(err)
		return -1, err
	}
	time := NewTime()
	result, err := stmtOut.Exec(9, time.DeletedAt, topic.Id)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return rows, nil
}
