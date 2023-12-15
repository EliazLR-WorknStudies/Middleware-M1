import json
from flask import jsonify
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user, logout_user

from src.models.http_exceptions import *
import src.repositories.users as users_repository


songs_url = "http://localhost:8080/songs/"     # URL de l'API songs (golang)
ratings_url = "http://localhost:8082/ratings/" # URL de l'API ratings (golang)


def get_song(id):
   
    # Appel à l'API Go pour obtenir la liste des chansons
    songs_response = requests.request(method="GET", url=songs_url+id)
    song = songs_response.json()

    # Appel à l'API Go pour obtenir la liste des ratings
    ratings_response = requests.request(method="GET", url=ratings_url)
    ratings = ratings_response.json()

    song_ratings = unify_song_ratings(song, ratings)
    return song_ratings
    

def get_songs():
    # Appel à l'API Go pour obtenir la liste des chansons
    songs_response = requests.request(method="GET", url=songs_url)
    songs = songs_response.json()

    # Appel à l'API Go pour obtenir la liste des ratings
    ratings_response = requests.request(method="GET", url=ratings_url)
    ratings = ratings_response.json()

    songs_with_ratings = []
    for song in songs:
        song_ratings = unify_song_ratings(song, ratings)
        songs_with_ratings.append(song_ratings)
    return jsonify(songs_with_ratings)

def unify_song_ratings(song, ratings):
    song_with_rating = {
        "SongAuthor": song["songauthor"],
        "SongName": song["songname"],
        "SongGenre": song["songgenre"],
        "Id": song["id"],
        "ratings": []
    }

     # Ajouter les évaluations associées à la chanson
    for rating in ratings:
            if rating["idSong"] == song["id"]:
                song_with_rating["ratings"].append({
                    "id": rating["id"],
                    "Comment": rating["comment"],
                    "Rating": rating["rating"],
                    "IdSong": rating["idSong"],
                    "IdUser": rating["idUser"]
                })
    return song_with_rating