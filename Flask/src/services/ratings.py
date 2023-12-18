import json
from flask import jsonify
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user, logout_user
from Flask.src.schemas.song import SongSchema
from Flask.src.schemas.ratings import RatingSchema

from src.models.http_exceptions import *
import src.repositories.users as users_repository


ratings_url = "http://localhost:8082/" # URL de l'API ratings (golang)

def get_ratings_from_song(id_song):
    
    ratings_response = requests.request(method="GET", url=ratings_url+"song/"+id_song)
    ratings_from_song = ratings_response.json()

    return ratings_from_song, ratings_response.status_code

def get_rating(id_rating):
    ratings_response = requests.request(method="GET", url=ratings_url+"ratings/"+id_rating)
    rating_from_song = ratings_response.json()
    
    return rating_from_song, ratings_response.status_code

def add_rating(id_song, rating_scheme):
    # on récupère le schéma rating pour la requête vers l'API rating
    rating_schema = RatingSchema().loads(json.dumps(rating_scheme), unknown=EXCLUDE)
    rating_schema.idSong = id_song
    ratings_response = requests.request(method="POST", url=ratings_url+"ratings/", json=rating_schema)

    if ratings_response.status_code != 201:
        return ratings_response.json(), ratings_response.status_code

    
    return ratings_response.json(), ratings_response.status_code