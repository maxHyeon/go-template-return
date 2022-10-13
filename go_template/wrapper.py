from ctypes import *
import os

root_dir = os.path.dirname(__file__)
shared_lib = os.path.join(root_dir, 'bind', 'template.so')
lib = cdll.LoadLibrary(shared_lib)
#lib = CDLL(shared_lib)

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

def get_go_string(val):
    return GoString(val.encode('utf-8'), len(val))

def get_go_path(file):
    if not os.path.isabs(file) and file:
        file = os.path.join(os.getcwd(),file)
    return get_go_string(file)

def render_template(template, value_file, output):
    template = get_go_path(template)
    value_file = get_go_path(value_file)
    if output != 'return' :
        output = get_go_path(output)
    elif output == 'return' :
        output = get_go_string(output)

    lib.RenderTemplate.argtypes = [GoString, GoString, GoString]
    lib.RenderTemplate.restype = c_char_p 
    return lib.RenderTemplate(template, value_file, output).decode('utf-8')
    

#render_template('tests/sample.tmpl', 'tests/values.yml','')
