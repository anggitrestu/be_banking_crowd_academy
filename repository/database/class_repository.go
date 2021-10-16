package database

import (
	"banking_crowd/models/classes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassRepository interface {
	Save(class classes.Class) (classes.Class, error)
	FindByIdTutor(TutorID int) ([]classes.Class, error)
	FindAll() ([]classes.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *classRepository {
	return &classRepository{db}
}

func (r *classRepository) Save(class classes.Class) (classes.Class, error) {
	err := r.db.Create(&class).Error
	if err != nil {
		return class, err
	}

	return class, nil
}

func (r *classRepository) FindByIdTutor(TutorID int) ([]classes.Class, error) {
	var classes []classes.Class
	err := r.db.Where("tutor_id = ? ", TutorID).Find(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}

func (r *classRepository) FindAll() ([]classes.Class, error) {
	var classes []classes.Class
	// err := r.db.Joins("Tutor").Find(&classes).Error
	// err := r.db.Raw("SELECT * FROM classes INNER JOIN my_classes ON  my_classes.class_id = classes.id").Scan(&classes).Error
	err := r.db.Preload(clause.Associations).Find(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}

/*

SELECT * FROM classes INNER JOIN my_classes ON  my_classes.class_id = classes.id;


SELECT Orders.OrderID, Customers.CustomerName
FROM Orders
INNER JOIN Customers ON Orders.CustomerID = Customers.CustomerID;

SELECT Orders.OrderID, Customers.CustomerName, Shippers.ShipperName
FROM ((Orders
INNER JOIN Customers ON Orders.CustomerID = Customers.CustomerID)
INNER JOIN Shippers ON Orders.ShipperID = Shippers.ShipperID);
*/
