package media

//import "go.mongodb.org/mongo-driver/v2/bson"

type Movie struct {
	//ID bson.ObjectID `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	ReleaseDate string `bson:"release_date" json:"release_date"`
}

