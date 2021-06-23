package firebasedb

import "time"

const (
	// defaults inital data
	QUESTION_LEVELS_COLLECTION = "question_levels"
	//
	USERS_COLLECTION            = "users"
	LEARNING_PATHS_COLLECTION   = "learning_paths"
	SUBJECTS_COLLECTION         = "subjects"
	TOPICS_COLLECTION           = "topics"
	QUESTIONS_COLLECTION        = "questions"
	DELETION_LOGGERS_COLLECTION = "deletion_loggers"
)

type User struct {
	ID        string    `json:"id" firestore:"id"`
	Name      string    `json:"name" firestore:"name"`
	Email     string    `json:"email" firestore:"email"`
	Active    bool      `json:"active" firestore:"active"`
	FCMTokens []string  `firestore:"fcm_tokens"`
	CreatedAt time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

type Subject struct {
	ID               string    `json:"id" firestore:"id"`
	Name             string    `json:"name" firestore:"name"`
	Description      string    `json:"description" firestore:"description"`
	TopicIds         []string  `json:"topic_ids" firestore:"topic_ids"`
	QuestionLevelIds []string  `json:"question_level_ids" firestore:"question_level_ids"`
	CreatedAt        time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt        time.Time `firestore:"updated_at"`
}

type Topic struct {
	ID          string    `json:"id" firestore:"id"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	SubjectId   string    `json:"subject_id" firestore:"subject_id"`
	TheoryId    string    `json:"theory_id" firestore:"theory_id"`
	CreatedAt   time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

// Available levels are
// Basic - 1, Intermediate -3, Advanced - 5, Coding - 5 , Realtime - 5
type QuestionLevel struct {
	ID        string    `json:"id" firestore:"id"`
	Name      string    `json:"name" firestore:"name"`
	Score     int       `json:"score" firestore:"score"`
	CreatedAt time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// question, a,b,c,d  strings are markdown
type Question struct {
	ID          string    `json:"id" firestore:"id"`
	Question    string    `json:"question" firestore:"question"`
	OptionA     string    `json:"a" firestore:"a"`
	OptionB     string    `json:"b" firestore:"b"`
	OptionC     string    `json:"c" firestore:"c"`
	OptionD     string    `json:"d" firestore:"d"`
	Answer      string    `json:"answer" firestore:"answer"`
	Explanation string    `json:"explanation" firestore:"explanation"`
	SubjectId   string    `json:"subject_id" firestore:"subject_id"`
	TopicId     string    `json:"topic_id" firestore:"topic_id"`
	Level       string    `json:"level_id" firestore:"level_id"`
	CreatedAt   time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

type LearningPath struct {
	ID          string    `json:"id" firestore:"id"`
	Name        string    `json:"name" firestore:"name"`
	Icon        string    `json:"icon" firestore:"icon"`
	Description string    `json:"description" firestore:"description"`
	SubjectIds  []string  `json:"subject_ids" firestore:"subject_ids"`
	Published   bool      `json:"published" firestore:"published"`
	CreatedAt   time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

// Deletionlogger
//
// Allowed names are same as collection names
type DeletionLogger struct {
	ID        string    `json:"id" firestore:"id"`
	DeletedId string    `json:"deleted_id" firestore:"deleted_id"`
	Name      string    `json:"name" firestore:"name"`
	CreatedAt time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// Insights
// TODO Total Questions, Total Subjects, Total Topics
// TODO overall questions count ,overall subjects count, overall topics count
// TODO Total tests taken as of now
