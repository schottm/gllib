#version 330

out vec4 frag_colour;

uniform vec3 colour;

void main() {

    frag_colour = vec4(colour, 1);
}