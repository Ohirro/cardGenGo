# Fractal Generator Golang

This project generates fractal maps using the Diamond-Square algorithm. The generated maps can represent different types of terrains, such as mountains, plains, and water bodies.

---

## Examples

Here are some examples of generated fractal maps:

### Example 1: Warm Pastel
![Example 1](examples/warm_pastel.png)

---

### Example 2: CoolPastel
![Example 2](examples/cool_pastel1.png)

---

### Example 3: Cool Pastel 2
![Example 3](examples/cool_pastel2.png)

## Steps
---
### Running
![Example 1](examples/app_1.png)

---

### Extend Left
![Example 2](examples/app_2.png)

---

### Extend right
![Example 3](examples/app_3.png)

---

### Regenerate
![Example 3](examples/app_4.png)


## How to Use
0. **For Linux**:
    ```bash
    sudo apt update && sudo apt install -y libgl1-mesa-dev xorg-dev
    ```

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo-name/fractal-generator.git
2. Create a .env file in the root directory with the following content:

```FRACTAL_N=8
FRACTAL_K=0.5
RENDER_VIEW_ANGLE=100.0
RENDER_LIGHT_ANGLE=45.0
OUTPUT_FILE=fractal.png
```

_FRACTAL_N: Size of the map (2^N + 1).

FRACTAL_K: Roughness factor.

RENDER_VIEW_ANGLE: Viewing angle for perspective rendering.

RENDER_LIGHT_ANGLE: Light source angle for shadows.

OUTPUT_FILE: Name of the output file._

3. Build and run the application:

```go build
./fractal-generator
```

The generated fractal will be saved as a .png file in the root directory.
