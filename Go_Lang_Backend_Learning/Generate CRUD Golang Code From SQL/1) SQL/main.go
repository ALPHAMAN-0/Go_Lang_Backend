package main

import (
    "fmt"
    "log"
    "SQL/db"
    "SQL/models"
)

func main() {
    // Connect to DB
    db.Connect()

    // ──────────────────────────────
    // 👤 USER CRUD (name)
    // ──────────────────────────────

    // CREATE
    id, err := models.CreateUser("Charlie", "charlie@example.com")
    if err != nil {
        log.Fatal("Create failed:", err)
    }
    fmt.Printf("✅ Created user with ID: %d\n", id)

    // READ ALL
    users, err := models.GetAllUsers()
    if err != nil {
        log.Fatal("Read failed:", err)
    }
    fmt.Println("📋 All Users:")
    for _, u := range users {
        fmt.Printf("   ID: %d | Name: %s | Email: %s\n", u.ID, u.Name, u.Email)
    }

    // UPDATE name
    err = models.UpdateUserName(id, "Charlie Brown")
    if err != nil {
        log.Fatal("Update failed:", err)
    }
    fmt.Printf("✏️  Updated user %d name to 'Charlie Brown'\n", id)

    // ──────────────────────────────
    // 📱 PHONE CRUD (phone number)
    // ──────────────────────────────

    // CREATE
    phoneID, err := models.CreatePhone(id, "+1-555-9999")
    if err != nil {
        log.Fatal("Phone create failed:", err)
    }
    fmt.Printf("✅ Created phone with ID: %d\n", phoneID)

    // READ
    phone, err := models.GetPhoneByUserID(id)
    if err != nil {
        log.Fatal("Phone read failed:", err)
    }
    fmt.Printf("📱 Phone for user %d: %s\n", id, phone.PhoneNumber)

    // UPDATE phone number
    err = models.UpdatePhoneNumber(id, "+1-999-0000")
    if err != nil {
        log.Fatal("Phone update failed:", err)
    }
    fmt.Println("✏️  Updated phone number to '+1-999-0000'")

    // DELETE phone
    err = models.DeletePhone(id)
    if err != nil {
        log.Fatal("Phone delete failed:", err)
    }
    fmt.Println("🗑️  Deleted phone")

    // DELETE user
    err = models.DeleteUser(id)
    if err != nil {
        log.Fatal("Delete failed:", err)
    }
    fmt.Printf("🗑️  Deleted user %d\n", id)
}