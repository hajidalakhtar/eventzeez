package seeder

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func Seed(db *gorm.DB) {
	// Create users
	user1ID := uuid.New()
	CreateUser(db, user1ID, "user1", "user1@example.com", "user", "password1")
	user2ID := uuid.New()
	CreateUser(db, user2ID, "user2", "user2@example.com", "user", "password2")
	user3ID := uuid.New()
	CreateUser(db, user3ID, "admin", "admin@example.com", "admin", "password3")

	// Create events
	event1ID := uuid.New()
	CreateEvent(db, event1ID, "Event 1", time.Now().AddDate(0, 1, 0), "12:40", user1ID, "Jakarta", "Event 1 description", "Offline")
	event2ID := uuid.New()
	CreateEvent(db, event2ID, "Event 2", time.Now().AddDate(0, 2, 0), "12:40", user2ID, "Jakarta", "Event 2 description", "Offline")

	// Create participants
	participant1ID := uuid.New()
	CreateParticipant(db, participant1ID, "Participant 1", "participant1@example.com", "0852155433613")
	participant2ID := uuid.New()
	CreateParticipant(db, participant2ID, "Participant 2", "participant2@example.com", "08521355433613")
	participant3ID := uuid.New()
	CreateParticipant(db, participant3ID, "Participant 3", "participant3@example.com", "0853155433613")

	//// Create participant events
	CreateParticipantEvent(db, participant1ID, event1ID)
	CreateParticipantEvent(db, participant2ID, event1ID)
	CreateParticipantEvent(db, participant3ID, event1ID)
	CreateParticipantEvent(db, participant2ID, event2ID)

	//// Create event users
	CreateEventUser(db, event1ID, user1ID)
	CreateEventUser(db, event2ID, user2ID)

	// Create tasks
	task1ID := uuid.New()
	CreateTask(db, task1ID, event1ID, "Task 1", false, "Task 1 Description", time.Now(), time.Now().AddDate(0, 0, 7))
	task2ID := uuid.New()
	CreateTask(db, task2ID, event2ID, "Task 2", false, "Task 2 Description", time.Now(), time.Now().AddDate(0, 0, 14))

	// Create task categories
	category1ID := uuid.New()
	CreateTaskCategory(db, category1ID, "Category 1", "#FF0000")
	category2ID := uuid.New()
	CreateTaskCategory(db, category2ID, "Category 2", "#00FF00")

	// Assign task categories
	CreateTaskCategoryAssignment(db, uuid.New(), task1ID, category1ID)
	CreateTaskCategoryAssignment(db, uuid.New(), task1ID, category2ID)
	CreateTaskCategoryAssignment(db, uuid.New(), task2ID, category1ID)

	// Create budgets
	budgetID1 := uuid.New()
	CreateBudget(db, budgetID1, event1ID, "Untuk Vendor", 1000000, "TES")
	budgetID2 := uuid.New()
	CreateBudget(db, budgetID2, event1ID, "Untuk Dekorasi", 1000000, "TES")

	// Create rundown
	rundownID1 := uuid.New()
	CreateRundown(db, rundownID1, "Rundown 1", "Rundown 1 Description", event1ID, time.Now(), time.Now().AddDate(0, 0, 7))
	rundownID2 := uuid.New()
	CreateRundown(db, rundownID2, "Rundown 2", "Rundown 2 Description", event1ID, time.Now(), time.Now().AddDate(0, 0, 7))
}
