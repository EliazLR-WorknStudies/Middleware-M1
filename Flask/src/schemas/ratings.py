from marshmallow import Schema, fields, validates_schema, ValidationError




class RatingSchema(Schema):
    id = fields.String(description="avis UUID")
    idSong = fields.String(description="musique UUID")
    idUser = fields.String(description="Utilisateur UUID")
    comment = fields.String(description="commentaire")
    rating = fields.String(description="note")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("idSong") or obj.get("idSong") == "") and \
               (not obj.get("idUser") or obj.get("idUser") == "") and \
               (not obj.get("comment") or obj.get("comment") == "") and \
               (not obj.get("rating") or obj.get("rating") == "")

    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "id" not in data or data["id"] == "" or \
           "idSong" not in data or data["idSong"] == "" or \
           "idUser" not in data or data["idUser"] == "" or \
           "comment" not in data or data["comment"] == "" or \
           "rating" not in data or data["rating"] == "":
            raise ValidationError("[id ,idSong, idUser, comment, rating] must all be specified")
        if data["rating"] > 10:
            data["rating"] = 10
            