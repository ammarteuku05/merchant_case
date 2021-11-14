# merchant_test_case

Before run This project : 

```
go mod tidy
```

For run :
```
go run . 
```
## Entity Relationship Diagram
[Entity Relationship Diagram](https://drive.google.com/file/d/11jhS05t-C-Xa7a81fDyFiR3vZbSzsmzU/view?usp=sharing 'ERD')


## Data Manipulation Language 
### Table Users

**Get all users**

   ```
SELECT id, full_name , email
FROM users
ORDER BY created_at DESC
   ```

**Register User**

   ```
INSERT INTO users(id, full_name, email, password, role, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

**Login User**

   ```
SELECT * 
FROM users 
WHERE email = ?
   ```

**Find User By Id**

   ```
SELECT * 
FROM users 
WHERE id = ?
   ```

**Delete User By Id**

   ```
DELETE FROM users 
WHERE id = ?
   ```

**Update User By Id**

   ```
UPDATE users 
SET full_name= ?, email= ?, password= ?,updated_at = ?
WHERE id = ?
   ```

**Find Outlet User By Id**

   ```
SELECT * FROM outlets 
WHERE id = ?
   ```

**Create Outlet by User**

   ```
INSERT INTO outlets(id, outlet_name, picture, user_id, created_at, updated_at)
VALUES(?,?,?,?,?,?)
   ```

**Get All Outlets**

   ```
SELECT id, outlet_name, picture, user_id
FROM outlets
ORDER BY created_at DESC
   ```



### Table Products

**Create Product**

   ```
INSERT INTO products(id, product_name, price, sku, picture, created_at, updated_at, outlet_id)
VALUES(?,?,?,?,?,?,?,?)
   ```

**Create Image**

   ```
INSERT INTO image_products(id, display_image, product_id)
VALUES(?,?,?)
   ```

**Find All Product**

   ```
SELECT id, product_name, price, sku, picture, outlet_id
FROM products
ORDER BY created_at DESC
   ```

**Find Product By Id**

   ```
SELECT id, product_name, price, sku, picture, outlet_id 
FROM products 
WHERE id = ?
   ```

**Update Product By Id**

   ```
UPDATE products 
SET product_name = ?, price = ?, sku = ?, picture = ?, outlet_id = ?, updated_at = ?
WHERE id = ?
   ```

**Delete Product By Id**

   ```
DELETE FROM products 
WHERE id = ?
   ```

## Activity Diagram 
[Activity Diagram](https://drive.google.com/file/d/1cJUsI76H3o-fRoO15ilZcqjETlvPMy8q/view?usp=sharing 'Activiy Digram')


## Use Case Diagram 
[Use Case Diagram](https://drive.google.com/file/d/1H2KkTsbghGbGN0DAIr3CGQTeAHbG0sFE/view?usp=sharing 'Use Case')

### add Information
for doc api you can use in posmant/merchant.postman_collection.json
