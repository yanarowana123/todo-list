package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo/internal/entity"
	"todo/internal/repository/mongodb/dto"
)

type Repository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return Repository{db, db.Collection("tasks")}
}

func (r Repository) Create(ctx context.Context, task entity.Task) (string, error) {
	t := dto.CreateTask{Title: task.Title, Status: int(task.Status), ActiveAt: task.ActiveAt, CreatedAt: task.CreatedAt}
	res, err := r.collection.InsertOne(ctx, t)

	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (r Repository) Update(ctx context.Context, task entity.Task) error {
	update := bson.D{{"$set", bson.D{{"title", task.Title},
		{"activeAt", task.ActiveAt}, {"status", task.Status}}}}

	objectId, err := primitive.ObjectIDFromHex(task.Id)

	if err != nil {
		return nil
	}

	_, err = r.collection.UpdateByID(ctx, objectId, update)
	return err
}

func (r Repository) Find(ctx context.Context, id string) *entity.Task {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": objectId})

	repoTask := dto.Task{}
	err = result.Decode(&repoTask)

	if err != nil {
		return nil
	}

	task := entity.Task{repoTask.Id, repoTask.Title, entity.Status(repoTask.Status), repoTask.ActiveAt, repoTask.CreatedAt}
	return &task
}

func (r Repository) Delete(ctx context.Context, task entity.Task) error {
	objectId, err := primitive.ObjectIDFromHex(task.Id)

	if err != nil {
		return err
	}

	r.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return nil
}

func (r Repository) List(ctx context.Context, status entity.Status) ([]entity.Task, error) {
	var tasks []entity.Task

	today := time.Now().UTC()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)

	var filteringCriteria map[string]interface{}
	if status == entity.BrandNew {
		filteringCriteria = bson.M{
			"$match": bson.M{
				"status":   status,
				"activeAt": bson.M{"$lte": today},
			},
		}
	} else {
		filteringCriteria = bson.M{
			"$match": bson.M{
				"status": status,
			},
		}
	}

	sortingCriteria := bson.M{
		"$sort": bson.M{
			"createdAt": -1,
		},
	}

	pipeline := bson.A{
		filteringCriteria,
		sortingCriteria,
	}

	cursor, err := r.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return tasks, err
	}
	defer cursor.Close(ctx)

	var repoTasks []dto.Task
	if err = cursor.All(context.Background(), &repoTasks); err != nil {
		return tasks, err
	}

	for _, repoTask := range repoTasks {
		tasks = append(tasks, entity.Task{repoTask.Id, repoTask.Title, entity.Status(repoTask.Status), repoTask.ActiveAt, repoTask.CreatedAt})
	}
	return tasks, nil
}
