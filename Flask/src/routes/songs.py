import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError
from Flask.src.schemas.user import UserUpdateSchema

from src.models.http_exceptions import *
from src.schemas.errors import *
import src.services.songs as songs_service

# from routes import users
songs = Blueprint(name="songs", import_name=__name__)

@songs.route('/', methods=['GET'])
@login_required
def get_songs():
    """
    ---
    get:
      description: Getting songs
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - songs
    """

    return songs_service.get_songs()

@songs.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """

    return songs_service.get_song(id)

@songs.route('/<id>/ratings', methods=['GET'])
# @login_required
def get_ratings(id):
    """
    ---
    get:
      description: Getting ratings of a song

      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    
    return songs_service.get_ratings(id)

@songs.route('/<id_song>/ratings/<id_rating>', methods=['GET'])
# @login_required
def get_rating(id_song, id_rating):
    """
    ---
    get:
      description: Getting ratings of a song

      parameters:
        - in: path
          name: id_song
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
          name: id_rating
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    try:
        return songs_service.get_rating(id_song, id_rating)
    except NotFound: 
          error = NotFoundSchema().loads(json.dumps({"message": "Can't manage other users"}))
          return error, error.get("code")
    except Forbidden: 
          error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other users"}))
          return error, error.get("code")
    
  
