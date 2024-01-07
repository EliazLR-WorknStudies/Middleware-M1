from flask import request
import yaml


def contentNegociation(data,err):
    #Renvoie les erreurs rencontrés si elles existent
    if((err<200)or(err>299)):
        return data, err
    
    #Recupère le type demandé dans la requete
    acceptHeader = request.headers.get('Accept', '').lower()
    
    #Si il demande du YAML on transforeme la donnée
    if(acceptHeader=='application/yaml'):
        return yaml.dump(data),err
    
    #Si il demande du JSON on renvoie tel quel
    if(acceptHeader=='application/json'):
        return data,err
    
    # Si il demande un truc inconnu on renvoie une 406
    return data, 406
