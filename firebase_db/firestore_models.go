package firebasedb

import "time"

const (
	// defaults inital data
	QUESTION_LEVELS_COLLECTION     = "question_levels"
	PRACTICE_TEST_TYPES_COLLECTION = "practice_test_types"
	//
	USERS_COLLECTION                      = "users"
	USER_PURCHASED_PRODUCTS_COLLECTION    = "user_purchased_products"
	LEARNING_PATHS_COLLECTION             = "learning_paths"
	SUBJECTS_COLLECTION                   = "subjects"
	SUBJECTS_QUESTIONS_MAPPING_COLLECTION = "subject_questions_mappings"
	TOPICS_COLLECTION                     = "topics"
	THEORIES_COLLECTION                   = "theories"
	QUESTIONS_COLLECTION                  = "questions"
	PRACTICE_TESTS_COLLECTION             = "practice_tests"
	DELETION_LOGGERS_COLLECTION           = "deletion_loggers"
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

// public_content, premium_content, additional_content strings are html file urls
type Theory struct {
	ID         string    `json:"id" firestore:"id"`
	Public     []string  `json:"public" firestore:"public"`
	Premium    []string  `json:"premium" firestore:"premium"`
	Additional []string  `json:"additional" firestore:"additional"`
	SubjectId  string    `json:"subject_id" firestore:"subject_id"`
	TopicId    string    `json:"topic_id" firestore:"topic_id"`
	CreatedAt  time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt  time.Time `firestore:"updated_at"`
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
	ID        string    `json:"id" firestore:"id"`
	Question  string    `json:"question" firestore:"question"`
	OptionA   string    `json:"a" firestore:"a"`
	OptionB   string    `json:"b" firestore:"b"`
	OptionC   string    `json:"c" firestore:"c"`
	OptionD   string    `json:"d" firestore:"d"`
	Answer    string    `json:"answer" firestore:"answer"`
	SubjectId string    `json:"subject_id" firestore:"subject_id"`
	TopicId   string    `json:"topic_id" firestore:"topic_id"`
	Level     string    `json:"level_id" firestore:"level_id"`
	CreatedAt time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

type LearningPath struct {
	ID          string    `json:"id" firestore:"id"`
	Name        string    `json:"name" firestore:"name"`
	Description string    `json:"description" firestore:"description"`
	SubjectIds  []string  `json:"subject_ids" firestore:"subject_ids"`
	Published   bool      `json:"published" firestore:"published"`
	CreatedAt   time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}

// Available test types are
// Learning Path or Subject
type PracticeTestType struct {
	ID        string    `json:"id" firestore:"id"`
	Name      string    `json:"name" firestore:"name"`
	CreatedAt time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// key of question ids is going to be question level id and value is array of question ids
type SubjectQuestionMappings struct {
	ID          string              `json:"id" firestore:"id"`
	QuestionIds map[string][]string `json:"question_ids" firestore:"question_ids"`
	CreatedAt   time.Time           `firestore:"created_at,serverTimestamp"`
	UpdatedAt   time.Time           `firestore:"updated_at"`
}

// Test_status = ('started' , 'completed' )
type PracticeTest struct {
	ID            string            `json:"id" firestore:"id"`
	Email         string            `json:"email" firestore:"email"`
	TestType      string            `json:"test_type" firestore:"test_type"`
	TestStatus    string            `json:"test_status" firestore:"test_status"`
	QuestionIdMap map[string]string `json:"questions_ids_map" firestore:"questions_ids_map"`
	TotalMarks    int               `json:"total_marks" firestore:"total_marks"`
	ObtainedMarks int               `json:"obtained_marks" firestore:"obtained_marks"`
	CreatedAt     time.Time         `firestore:"created_at,serverTimestamp"`
	UpdatedAt     time.Time         `firestore:"updated_at"`
}

// User Purchased products
type UserPurchasedProduct struct {
	ID                string    `json:"id" firestore:"id"`
	Email             string    `json:"email" firestore:"email"`
	PurchasedDomainId string    `json:"purchased_domain_id" firestore:"purchased_domain_id"`
	CreatedAt         time.Time `firestore:"created_at,serverTimestamp"`
	UpdatedAt         time.Time `firestore:"updated_at"`
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
