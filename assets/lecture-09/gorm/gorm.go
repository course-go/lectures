package main

// OPEN START OMIT

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price float64 `gorm:"default:0.0"`
}

func main() {
	databaseURI := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(databaseURI))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	// OPEN END OMIT
	// CREATE START OMIT

	product := Product{Name: "Gopher", Price: 21.99}
	result := db.Create(&product)
	fmt.Println(product.ID) // backfills primary key values

	products := []*Product{
		&Product{Name: "Gopher", Price: 21.99},
		&Product{Name: "Hamster", Price: 15.49},
	}
	result = db.Create(products)

	product.Name = "Beaver"
	// INSERT INTO `products` (`name`) VALUES ("Beaver")
	db.Select("Name").Create(&product)
	db.Omit("Price").Create(&product)

	db.Model(&Product{}).Create(map[string]any{
		"Name": "Otter", "Price": 29.99,
	})

	// CREATE END OMIT
	// READ START OMIT

	// SELECT * FROM products ORDER BY id LIMIT 1;
	db.First(&product)

	// SELECT * FROM products LIMIT 1;
	db.Take(&product)

	// SELECT * FROM products ORDER BY id DESC LIMIT 1;
	result = db.Last(&product)
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// SELECT * FROM products;
	db.Find(&products)

	// SELECT * FROM products WHERE id = 10;
	db.Find(&product, 3)
	// SELECT * FROM users WHERE id IN (1,2,3);
	db.Find(&products, []int{1, 2, 3})

	// READ END OMIT
	// READ-CONDITIONS START OMIT

	// SELECT * FROM products WHERE id = 'some_uuid';
	db.First(&product, "id = ?", "some_uuid")

	// SELECT * FROM products WHERE name = 'Hamster';
	db.Find(&product, "name = ?", "Hamster")

	// SELECT * FROM products WHERE name <> 'Otter' AND price > 20.0;
	db.Find(&products, "name <> ? AND price > ?", "Otter", 20.0)

	// SELECT * FROM products WHERE name = 'Beaver';
	db.Find(&products, User{Name: "Beaver"})
	db.Find(&products, map[string]interface{}{"name": "Beaver"})

	// SELECT * FROM products WHERE NOT name = 'Gopher' ORDER BY id LIMIT 1;
	db.Not("name = ?", "Gopher").First(&product)

	// SELECT * FROM products OFFSET 5 LIMIT 10;
	db.Limit(10).Offset(5).Find(&products)

	// READ-CONDITIONS END OMIT
	// UPDATE START OMIT

	db.First(&product)
	product.Name = "Racoon"
	product.Price = 99.99
	// UPDATE products SET name="Racoon", price=99.99, updated_at="2024-04-21 03:24:10" WHERE id=1;
	db.Save(&product)

	// INSERT INTO products (`name`,`price`,`created_at`) VALUES ("T-shirt",19.99,"2024-04-21 02:55:22")
	db.Save(&Product{Name: "T-shirt", Price: 19.99})

	// UPDATE products SET name="Hoodie", price=44.49, updated_at="2024-04-21 03:28:42" WHERE id = 1
	db.Save(&Product{Model: gorm.Model{ID: 1}, Name: "Hoodie", Price: 44.49})

	// UPDATE products SET name='Beaver plush toy' WHERE name='Beaver';
	db.Model(User{}).Where("name = ?", "Beaver").Updates(Product{Name: "Beaver plush toy"})

	// UPDATE products SET name='Beanie', updated_at = '2024-04-21 03:42:36' WHERE id = 11;
	db.Model(&product).Updates(Product{Model: gorm.Model{ID: 11}, Name: "Beanie"})

	// UPDATE END OMIT
	// DELETE START OMIT

	// DELETE FROM products WHERE id = 6;
	db.Delete(&Product{}, 6)

	// DELETE FROM products WHERE id IN (2,3,5);
	db.Delete(&Product{}, []int{2, 3, 5})

	// DELETE FROM products WHERE name LIKE "%boo%";
	db.Where("name LIKE ?", "%boo%").Delete(&Product{})

	// DELETE FROM products WHERE price > 100.0;
	db.Delete(&Product{}, "price > ?", 100.0)

	products = []*Product{&Product{Model: gorm.Model{ID: 1}}, &Product{Model: gorm.Model{ID: 2}}}
	// DELETE FROM products WHERE id IN (1,2);
	db.Delete(&products)

	// DELETE FROM users WHERE name LIKE "%boo%" AND id IN (1,2);
	db.Delete(&products, "name LIKE ?", "%boo%")

	// DELETE END OMIT
	// SOFT-DELETE START OMIT

	product = Product{Model: gorm.Model{ID: 11}}
	// UPDATE products SET deleted_at="2024-04-21 03:58:11" WHERE id = 11;
	db.Delete(&product)

	// UPDATE products SET deleted_at="2024-04-21 03:58:28" WHERE age = 20;
	db.Where("name = ?", "Gopher").Delete(&User{})

	// SELECT * FROM products WHERE name='Gopher' AND deleted_at IS NULL;
	db.Where("name = Gopher").Find(&product)

	// SOFT-DELETE END OMIT
	// CONVENTIONS START OMIT

	db.Create(&product) // Sets `CreatedAt` to current time

	product = Product{
		Name: "Mug",
		Model: gorm.Model{
			CreatedAt: time.Now(),
		},
	}
	db.Create(&product) // `CreatedAt` won't be changed

	// To change its value, you could use `Update`
	db.Model(&product).Update("CreatedAt", time.Now())

	// CONVENTIONS END OMIT
	// RAW START OMIT

	var total int
	db.Raw("SELECT SUM(price) FROM products").Scan(&total)

	db.Raw("UPDATE products SET name = ? WHERE price = ? RETURNING id, name",
		"Overpriced toy", 100.00,
	).Scan(&products)

	db.Exec("UPDATE products SET deleted = ? WHERE id IN ?",
		time.Now(), []int64{1, 2, 3},
	)

	db.Exec("UPDATE products SET price = ? WHERE name = ?",
		gorm.Expr("price * ? + ?", 2, 5),
		"Gopher",
	)

	db.Exec("DROP TABLE products")

	// RAW END OMIT
}
