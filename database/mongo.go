package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// Ancienne méthode de connexion
func Connect(uri string, dbName string) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("❌ Connexion Mongo échouée :", err)
    }

    // Ping pour vérifier la connexion
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("❌ MongoDB ping échoué :", err)
    }

    log.Println("✅ Connexion à MongoDB réussie")
    MongoClient = client
    MongoDB = client.Database("goodWave")
}

// Nouvelle méthode selon la documentation MongoDB
func ConnectWithOptions(uri string, dbName string, serverAPI *options.ServerAPIOptions) error {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    // Configuration du client avec l'URI et les options du serveur API
    clientOptions := options.Client().
        ApplyURI(uri).
        SetServerAPIOptions(serverAPI)

    // Connexion au serveur
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Println("❌ Connexion Mongo échouée :", err)
        return err
    }

    // Ping pour vérifier la connexion
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        log.Println("❌ MongoDB ping échoué :", err)
        return err
    }

    log.Println("✅ Connexion à MongoDB réussie")
    MongoClient = client
    MongoDB = client.Database("goodWave")
    return nil
}