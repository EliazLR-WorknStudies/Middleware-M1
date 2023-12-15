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
def get_users():
    return songs_service.get_songs()

@songs.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
   
    
    return songs_service.get_song(id)
