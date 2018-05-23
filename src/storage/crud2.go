package storage

import (
	"fmt"
	"errors"
)

type Topic struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int 
	Content string
	Author string
	Topic *Topic
}
 
func (comment *Comment) Create() (err error){
	if comment.Topic == nil {
		err = errors.New("Post not found")
		return
	}
	query := `insert into comments (content, author, topic_id)
	values($1, $2, $3) returning id`
	err = Db.QueryRow(query, comment.Content,comment.Author, comment.Topic.Id).Scan(&comment.Id)
	return
}

func (topic *Topic) Create() (err error) {
	
	query :=`insert into topic( content, author)
	values($1, $2) returning id`
	err = Db.QueryRow(query, topic.Content, topic.Author).Scan(&topic.Id)
	return 
}

func GetTopic(id int) (topic Topic, err error){
	topic = Topic{}
	topic.Comments = []Comment{}

	query := "SELECT id, content, author from topic where id = $1"
	err = Db.QueryRow(query, id).Scan(&topic.Id, &topic.Content, &topic.Author)
	if err != nil {
		panic(err)
		return
	}

	secondQuery := "select id, content, author from comments where topic_id = $1"
	rows, err := Db.Query(secondQuery, id)

	if err != nil {
		panic(err)
		return 
	}

	for rows.Next() {
		comment:= Comment{Topic:&topic}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			panic(err)
			return
		}
		topic.Comments = append(topic.Comments, comment)

	}
	rows.Close()
	return
}

func CRUD2(){
	topic  := Topic{Content: "Hello World!", Author:"Sau Sheong"}
	err := topic.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(topic)

	comment := Comment{Content: "Good topic", Author: "Joe", Topic: &topic}
	err = comment.Create()
	if err != nil {
		panic(err)
	}
	readPost, _ := GetTopic(topic.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Topic)
}
