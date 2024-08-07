package repository

import (
	"context"
	"fmt"

	"dev.sachinborse/rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	MongoDBCollection *mongo.Collection
}

func (e *EmployeeRepository) insertEmployee(emp *models.Employee) (interface{}, error) {
	result, err := e.MongoDBCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (e *EmployeeRepository) findEmployeeByID(empID string) (*models.Employee, error) {
	var emp models.Employee

	if err := e.MongoDBCollection.FindOne(context.Background(), bson.D{{Key: "id", Value: empID}}).Decode(&emp); err != nil {
		return nil, err
	}

	return &emp, nil
}

func (e *EmployeeRepository) findAllEmployee() ([]models.Employee, error) {
	result, err := e.MongoDBCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []models.Employee
	err = result.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}
	return emps, nil
}

func (e *EmployeeRepository) updateEmployeeByID(empID string, updatedEmp *models.Employee) (int64, error) {
	result, err := e.MongoDBCollection.UpdateOne(context.Background(),
		bson.D{{Key: "id", Value: empID}},
		bson.D{{Key: "$set", Value: updatedEmp}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (e *EmployeeRepository) deleteEmployeeByID(empID string) (int64, error) {
	result, err := e.MongoDBCollection.DeleteOne(context.Background(),
		bson.D{{Key: "id", Value: empID}})

	if err != nil {
		return 0, nil
	}

	return result.DeletedCount, nil
}

func (e *EmployeeRepository) deleteAllEmployee() (int64, error) {
	result, err := e.MongoDBCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
