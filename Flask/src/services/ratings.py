import json
from flask import jsonify
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user, logout_user
from Flask.src.schemas.song import SongSchema

from src.models.http_exceptions import *
import src.repositories.users as users_repository


ratings_url = "http://localhost:8082/" # URL de l'API ratings (golang)
