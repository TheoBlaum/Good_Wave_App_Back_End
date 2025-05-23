#!/bin/bash

# Compiler et exécuter le script de conversion des IDs
echo "Conversion des IDs..."
cd convert_ids
go build
./convert_ids

# Compiler et exécuter le script d'import des données
echo "Import des données..."
cd ../import_data
go build
./import_data

echo "Processus terminé avec succès" 