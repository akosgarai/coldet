# Coldet

Collision detection library written in golang.

## Inspirations

[Learnopengl](https://learnopengl.com/In-Practice/2D-Game/Collisions/Collision-detection) tutorial about the 2D collision detection.
[Nehe](http://nehe.gamedev.net/tutorial/collision_detection/17005/) tutorial about the collision detection.
[MDN](https://developer.mozilla.org/en-US/docs/Games/Techniques/3D_collision_detection) tutorial about the 3D collision detection.

## Options

- AABB - AABB
- Point - Sphere
- Point - AABB
- AABB - Sphere
- Sphere - Sphere

## What is a Point?

It is a structure that holds only a position (x, y, z coordinates).

## What is an AABB

Axis aligned bounding box. It holds its center position, and the lenght of the sides of the box (x, y, z coordinates, width, length, height).

## What is a Sphere

It holds it's center point's position and the radius of the sphere (x, y, z coordinates, radius).
