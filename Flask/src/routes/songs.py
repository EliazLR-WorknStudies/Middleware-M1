import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError
from Flask.src.schemas.song import SongSchema
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

@songs.route('/', methods=['POST'])
# @login_required
def add_song():
    """
    ---
    post:
      description: Adding a song
      requestBody:
        required: true
        content:
            application/json:
                schema: SongSchema
      responses:
        '201':
          description: Created
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
    """

    # parser le body
    try:
        song_scheme = SongSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # enregistrer la musique
    try:
        return songs_service.add_song(song_scheme)
    except SomethingWentWrong:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    
@songs.route('/<id>', methods=['DELETE'])

#@login_required
def delete_song(id):
  """
  ---
  delete:
    description: Delete a user
    responses:
      '204':
        description: No centent
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
        - users
    """

    
  try:
      return songs_service.delete_song(id)
  except UnprocessableEntity:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
      return error, error.get("code")
  except Forbidden:
      error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
      return error, error.get("code")
  except Exception:
      error = SomethingWentWrongSchema().loads("{}")
      return error, error.get("code")


@songs.route('/<id>', methods=['PUT'])

#@login_required
def update_song(id):
  """
  ---
  delete:
    description: Update a user
    responses:
      '204':
        description: No centent
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
        - users
    """

  # parser le body
  try:
      song_scheme = SongSchema().loads(json_data=request.data.decode('utf-8'))
  except ValidationError as e:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
      return error, error.get("code")

  # modification de l'utilisateur (username, nom, mot de passe, etc.)
    
  try:
      return songs_service.update_song(song_scheme, id)
  except UnprocessableEntity:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
      return error, error.get("code")
  except Forbidden:
      error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
      return error, error.get("code")
  except Exception:
      error = SomethingWentWrongSchema().loads("{}")
      return error, error.get("code")






















## RATINGS SPACE



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
    
  
