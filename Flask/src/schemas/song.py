from marshmallow import Schema, fields, validates_schema, ValidationError




class SongSchema(Schema):
    songauthor = fields.String(description="auteur")
    songname = fields.String(description="nom")
    songgenre = fields.String(description="genre")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("songauthor") or obj.get("songauthor") == "") and \
               (not obj.get("songname") or obj.get("songname") == "") and \
               (not obj.get("songgenre") or obj.get("songgenre") == "")

    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "songname" not in data or data["songname"] == "" or \
           "songauthor" not in data or data["songauthor"] == "" or \
           "songgenre" not in data or data["songgenre"] == "":
            raise ValidationError("['songname','songauthor', songgenre] must all be specified")
