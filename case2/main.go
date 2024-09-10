package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const rawComments = `[
	{
		"commentId": 1,
		"commentContent": "Hai",
		"replies": [
			{
				"commentId": 11,
				"commentContent": "Hai juga",
				"replies": [
					{
						"commentId": 111,
						"commentContent": "Haai juga hai jugaa"
					},
					{
						"commentId": 112,
						"commentContent": "Haai juga hai jugaa"
					}
				]
			},
			{
				"commentId": 12,
				"commentContent": "Hai juga",
				"replies": [
					{
						"commentId": 121,
						"commentContent": "Haai juga hai jugaa"
					}
				]
			}
		]
	},
	{
		"commentId": 2,
		"commentContent": "Halooo"
	}
]`

type (
	Comment struct {
		CommentID      int       `json:"commentId"`
		CommentContent string    `json:"commentContent"`
		Replies        []Comment `json:"replies,omitempty"`
	}

	Comments []Comment
)

func calculateComment(comments Comments) (res int) {
	for _, comment := range comments {
		res++
		res += calculateComment(comment.Replies)
	}
	return
}

func main() {
	var comments Comments
	err := json.Unmarshal([]byte(rawComments), &comments)
	if err != nil {
		log.Fatal(err)
	}

	// using recursive to sum up the comments
	totalComment := calculateComment(comments)
	fmt.Println("===== nomor 5 =====")
	fmt.Println("total comment(s):", totalComment)
}
