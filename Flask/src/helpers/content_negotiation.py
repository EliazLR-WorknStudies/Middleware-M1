from flask import request
import yaml

def content_negotiation(data):
    # takes a tuple[Any, int] composed of the response body (json) and the status code, 
    accept_header = request.headers.get('Accept', '*/*')
    types = accept_header.split(',')

    # Parse the types into a list of (type, qvalue) tuples
    types = [(t.split(';')[0].strip(), float(t.split('q=')[1]) if 'q=' in t else 1) for t in types]

    # Sort the types by qvalue
    types.sort(key=lambda x: -x[1])

    # Iterate over the sorted types
    for content_type, _ in types:
        if content_type in ['application/json', '*/*']:
            return data[0], data[1]
        elif content_type == 'application/yaml':
            return yaml.dump(data[0]), data[1]

    # If no acceptable types were found, return a 406 Not Acceptable
    return 'Not Acceptable', 406
