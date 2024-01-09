import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError
from Flask.src.schemas.song import SongSchema
from Flask.src.schemas.user import UserUpdateSchema
from Flask.src.schemas.ratings import RatingSchema

from src.models.http_exceptions import *
from src.schemas.errors import *
from src.helpers.content_negotiation import *
import src.services.songs as songs_service
import src.services.ratings as ratings_service

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

    response,err= songs_service.get_songs()
    return contentNegociation(response,err)


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

    response,err= songs_service.get_song(id)
    return contentNegociation(response,err)

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
        response, err = songs_service.add_song(song_scheme)
    except SomethingWentWrong:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    
    return contentNegociation(response,err)
    
@songs.route('/<id>', methods=['DELETE'])

#@login_required
def delete_song(id):
  """
  ---
  delete:
    description: Delete a song
    responses:
      '204':
        description: No content
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

    
  try:
      response, err = songs_service.delete_song(id)
  except UnprocessableEntity:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
      return error, error.get("code")
  except Forbidden:
      error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
      return error, error.get("code")
  except Exception:
      error = SomethingWentWrongSchema().loads("{}")
      return error, error.get("code")
  
  return contentNegociation(response,err)


@songs.route('/<id>', methods=['PUT'])

#@login_required
def update_song(id):
  """
  ---
  put:
    description: Update a song
    responses:
      '204':
        description: No content
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

  # parser le body
  try:
      song_scheme = SongSchema().loads(json_data=request.data.decode('utf-8'))
  except ValidationError as e:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
      return error, error.get("code")

  # modification de l'utilisateur (username, nom, mot de passe, etc.)
    
  try:
      response, err = songs_service.update_song(song_scheme, id)
  except UnprocessableEntity:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
      return error, error.get("code")
  except Forbidden:
      error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
      return error, error.get("code")
  except Exception:
      error = SomethingWentWrongSchema().loads("{}")
      return error, error.get("code")
  
  return contentNegociation(response,err)















## RATINGS SPACE



@songs.route('/<id_song>/ratings', methods=['GET'])
# @login_required
def get_ratings_from_song(id_song):
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
          - ratings
    """
    
    response, err = ratings_service.get_ratings(id_song)
    return contentNegociation(response,err)

@songs.route('/<id_song>/ratings/<id_rating>', methods=['GET'])
# @login_required
def get_rating(id_rating):
    """
    ---
    get:
      description: Getting ratings of a song

      parameters:
        - in: path
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
          - ratings
    """

    response, err = ratings_service.get_rating(id_rating)
    return contentNegociation(response,err)
    
    

@songs.route('/<id_song>/ratings', methods=['POST'])
# @login_required
def add_rating(id_song):
    """
    ---
    post:
      description: creating a new rating for a song

      parameters:
        - in: path
          name: id_song
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
          requestBody:
            required: true
            content:
              application/json:
                schema: Rating Schema
      responses:
        '201':
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
        '422':
          description: Unprocessable Entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Internal Server Error
          content:
            application/json:
              schema: InternalServerError
            application/yaml:
              schema: InternalServerError
      tags:
          - ratings
          - songs
    """
    
    try:
        rating_scheme = RatingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # enregistrer le commentaire
    try:
        response, err = ratings_service.add_rating(id_song, rating_scheme)
    except SomethingWentWrong:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    
    return contentNegociation(response,err)
  

@songs.route('/<id_song>/ratings/<id_rating>', methods=['PUT'])
# @login_required
def update_rating(id_rating):
    """
    ---
    put:
      description: Updating a rating of a song

      parameters:
        - in: path
          name: id_rating
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
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
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable Entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - ratings
          - songs
    """
    
    # parser le body
    try:
        rating_update = RatingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification du rating
    try:
        response,err = ratings_service.update_rating(id_rating, rating_update)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other user's ratings"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
    
    return contentNegociation(response,err)

    
@songs.route('/<id_song>/ratings/<id_rating>', methods=['DELETE'])
# @login_required
def delete_rating(id_rating):
    """
    ---
    delete:
      description: delete a rating of a song

      parameters:
        - in: path
          name: id_rating
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: no Content
          content:
            application/json:
              schema: NoContent
            application/yaml:
              schema: NoContent
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - ratings
          - songs
    """
    try:
        response,err = ratings_service.delete_rating(id_rating)
    except NotFound: 
          error = NotFoundSchema().loads(json.dumps({"message": "Can't manage other user's raitings"}))
          return error, error.get("code")
    except Forbidden: 
          error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other user's raitings"}))
          return error, error.get("code")
    
    return contentNegociation(response,err)
    
