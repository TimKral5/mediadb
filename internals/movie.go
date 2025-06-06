package internals

type Movie struct {
	Title string `bson:"title"`
	Description string `bson:"description"`
	ReleaseDate string `bson:"release_date"`
}

