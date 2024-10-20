package models

type YogurtTask struct {
  TaskID      string `json:"task_id"`
	Emoji       string `json:"emoji"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TaskSize    int    `json:"task_size"`
	Difficulty  int    `json:"difficulty"`
	Priority    int    `json:"priority"`
	CategoryID  int    `json:"category_id"`
}
