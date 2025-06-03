# 📞 Phone Number Normalizer (Go + SQLite)

This Go program demonstrates how to store, normalize, and manage phone numbers in a SQLite database. It handles phone numbers in various formats, removes duplicates, and ensures that all stored numbers contain only digits.

## 🧩 Features

- Stores phone numbers in a SQLite database.
- Supports phone numbers in various formats (e.g., `(123) 456-7890`, `123 456 7891`).
- Normalizes all phone numbers to digit-only format (e.g., `1234567890`).
- Removes duplicate numbers after normalization.

## 📦 Requirements

- Go 1.16+
- SQLite
- Go SQLite driver: [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3)

## 📁 Installation

1. **Clone the repository or copy the code:**

   ```bash
   git clone <your-repo-url>
   cd <repo-directory>
   ```

2. **Install dependencies:**

   ```bash
   go get github.com/mattn/go-sqlite3
   ```

3. **Run the program:**

   ```bash
   go run main.go
   ```

## ⚙️ How It Works

### 1. Setup

- Creates a SQLite database file `database.db`.
- Ensures the `phone_numbers` table exists.
- Clears any existing data.
- Inserts a mix of hardcoded and test phone numbers in various formats.

### 2. Normalization

- All phone numbers are normalized by stripping out non-digit characters using a regular expression.
- If two phone numbers normalize to the same value, the duplicate is deleted.

### 3. Output

- The program prints all phone numbers before and after normalization.

## 🧪 Example

### Input:

```
1234567890
123 456 7891
(123) 456 7892
...
```

### Output:

```
Before:
1234567890
123 456 7891
(123) 456 7892
...

After:
1234567890
1234567891
1234567892
...
```

## 🧼 Normalization Logic

The normalization is handled by this function:

```go
func remainDigitOnly(phone_number string) string {
	r, _ := regexp.Compile(`[^\d]`)
	return r.ReplaceAllString(phone_number, "")
}
```

It removes all non-digit characters from the input and duplicates are removed in the database.

## 🛠️ File Structure

```
main.go         # The main Go program
database.db     # Generated SQLite database
```
