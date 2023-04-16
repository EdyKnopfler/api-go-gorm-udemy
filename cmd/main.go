package main

import (
	"fmt"

	"com.derso/curso_creuto/gorm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Facilitando um pouco a vida aqui
	databaseName := "postgres"
	userName := "postgres"
	password := "mysecretpassword"
	host := "localhost"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		host, userName, password, databaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Evita erros no relacionamento bidirecional; criamos manualmente depois
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}

	// Criamos as tabelas
	// Recomendado o golang-migrations no lugar disto!
	db.Migrator().AutoMigrate(&model.Member{}, &model.Note{})

	// Agora criamos as FKs
	db.Migrator().CreateConstraint(&model.Member{}, "Notes")
	db.Migrator().CreateConstraint(&model.Member{}, "Connections")

	// Hora de brincar
	user1 := &model.Member{Name: "Kânia", Email: "kania@gato.com"}
	db.Create(&user1)
	user2 := &model.Member{Name: "Kurko", Email: "kurko@cachorro.com"}
	db.Create(&user2)

	// Ligações (atenção a onde requer ou não o ponteiro!)
	db.Model(&user2).Association("Connections").Append(user1)
	db.Save(user2)

	db.Model(&user1).Association("Notes").Append(&model.Note{
		Author: *user1, Text: "Nota louca"})
	db.Save(user1)

	// Carregando
	kania := &model.Member{}
	db.Preload("Notes").Preload("Connections").First(&kania, 1)
	fmt.Printf("Nota do Kânia: %s\n", kania.Notes[0].Text)
	fmt.Printf("Kânia tem %d conexões.\n", len(kania.Connections))

	kurko := &model.Member{}
	db.Preload("Notes").Preload("Connections").First(&kurko, 2)
	fmt.Printf("Kurko tem %d notas.\n", len(kurko.Notes))
	fmt.Printf("Kurko tem uma conexão com: %s\n", kurko.Connections[0].Name)
}
