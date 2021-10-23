
# Document Service

Document service propose un exemple de service RESTful, développé en Golang et possédant des tests unitaires.   L'API et les test unitaires sont développés grâce aux framework Gin et Testify. Ce service suit les contraintes d'architectures "clean architecture".

## Contexte
Le service permet d'interagir avec des Documents.
Un document un composé d'un ID unique de 24 caractères généré par le service à sa création, ainsi que d'un nom et d'une description.  
Le service permet de créer, récupérer ou supprimer des Documents.

## Mise en route
### Lancement
Pour lancer le programme, se placer dans le répertoire app/main et entrer la commande :
`go run .`
Le serveur se lance alors en local sur le port 8081

### Lancement (Tests)
Pour lancer tous les test du programme, se placer dans le répertoire source et entrer la commande :
`go test -v  ./...`

## Utilisation
- Création :
    - POST http://localhost:8081/document
    - `body : {
      "description" : "C'est un super document",
      "name" : "superdoc2000"
      }`

- Suppression :
    - DELETE http://localhost:8081/document/:id

- Récupération :
    -  GET http://localhost:8081/document/:id

## References
- github.com/gin-gonic/gin v1.7.4
- github.com/stretchr/testify v1.7.0


*Développé par Pierre-Jean Besnard*
