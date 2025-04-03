from PIL import Image
import numpy as np
import random

def create_tile(rgb_values, filename):
    pixels = np.array(rgb_values, dtype=np.uint8).reshape(16, 16, 3)
    image = Image.fromarray(pixels)
    image.save(f"{filename}.png")

# Color palette
DARK_GREEN = [34, 139, 34]
LIGHT_GREEN = [124, 252, 0]
BROWN = [139, 69, 19]
DARK_BROWN = [101, 67, 33]

def add_color_noise(base_color, intensity=10):
    """Add slight random variation to a color"""
    return [max(0, min(255, c + random.randint(-intensity, intensity))) for c in base_color]

def blend_colors(color1, color2, ratio):
    """Blend two colors together based on ratio (0.0 to 1.0)"""
    return [int(c1 * (1 - ratio) + c2 * ratio) for c1, c2 in zip(color1, color2)]

def generate_noise_map(threshold=0.5):
    """Generate a 16x16 noise map"""
    return [[random.random() for _ in range(16)] for _ in range(16)]

def generate_grass_tile():
    pixels = []
    # Generate noise maps for variation
    primary_noise = generate_noise_map(0.6)
    texture_noise = generate_noise_map(0.3)
    
    for y in range(16):
        for x in range(16):
            # Base color selection with primary noise
            base_color = LIGHT_GREEN if primary_noise[y][x] > 0.7 else DARK_GREEN
            
            # Add texture variation
            if texture_noise[y][x] > 0.8:
                # Occasional darker patches
                base_color = [max(0, c - 20) for c in base_color]
            elif texture_noise[y][x] < 0.2:
                # Occasional lighter patches
                base_color = [min(255, c + 20) for c in base_color]
            
            # Add subtle random noise
            final_color = add_color_noise(base_color, 5)
            pixels.extend(final_color)
    
    return pixels

def generate_dirt_tile():
    pixels = []
    # Generate multiple noise maps for different features
    primary_noise = generate_noise_map()
    detail_noise = generate_noise_map()
    
    for y in range(16):
        for x in range(16):
            # Base color selection
            noise_val = primary_noise[y][x]
            detail = detail_noise[y][x]
            
            # Create more natural dirt patterns
            if noise_val > 0.7:
                base_color = DARK_BROWN
            elif noise_val < 0.3:
                base_color = [c + 10 for c in BROWN]  # Slightly lighter brown
            else:
                base_color = BROWN
            
            # Add texture details
            if detail > 0.9:
                # Occasional darker spots (small rocks or shadows)
                base_color = [max(0, c - 25) for c in base_color]
            elif detail < 0.1:
                # Occasional lighter spots (small pebbles)
                base_color = [min(255, c + 15) for c in base_color]
            
            # Add subtle random noise
            final_color = add_color_noise(base_color, 3)
            pixels.extend(final_color)
    
    return pixels

def generate_mixed_tile():
    pixels = []
    # Generate noise maps for both grass and dirt
    transition_noise = generate_noise_map()
    detail_noise = generate_noise_map()
    
    for y in range(16):
        for x in range(16):
            # Create a noisy transition line
            base_threshold = y - x + random.randint(-2, 2)
            noise_influence = transition_noise[y][x] * 3
            
            # Determine if we're in the grass or dirt region
            if base_threshold + noise_influence > 2:
                # Grass region
                if detail_noise[y][x] > 0.7:
                    base_color = LIGHT_GREEN
                else:
                    base_color = DARK_GREEN
                final_color = add_color_noise(base_color, 5)
            else:
                # Dirt region
                if detail_noise[y][x] > 0.8:
                    base_color = DARK_BROWN
                else:
                    base_color = BROWN
                final_color = add_color_noise(base_color, 3)
                
            pixels.extend(final_color)
    
    return pixels

def generate_grass_with_flowers():
    # Colors
    FLOWER_YELLOW = [255, 255, 0]
    FLOWER_PURPLE = [128, 0, 128]
    FLOWER_CENTER = [255, 140, 0]
    
    # First generate the natural grass background using our improved grass generator
    pixels = generate_grass_tile()  # This gives us our noisy grass base
    
    # Convert the flat pixel list back to a 2D array for easier manipulation
    pixels_2d = [pixels[i:i+48] for i in range(0, len(pixels), 48)]  # 48 because each pixel is 3 values (RGB)
    
    # Flower positions (x, y coordinates)
    flowers_0 = [(4, 6), (12, 10), (4, 12)]
    flowers_1 = [(8, 12), (14, 7)]
    
    # Add flowers over the grass
    for y in range(16):
        for x in range(16):
            pixel_index = (y * 16 + x) * 3  # Calculate the index in our flat pixel list
            
            # If this is a flower center
            if (x, y) in flowers_0:
                pixels[pixel_index:pixel_index+3] = FLOWER_CENTER
            # If this is adjacent to a flower (for petals)
            elif any((abs(x - fx) == 1 and y == fy) or (x == fx and abs(y - fy) == 1) 
                    for fx, fy in flowers_0):
                # Add some variation to the flower petals
                petal_color = add_color_noise(FLOWER_YELLOW, 5)
                pixels[pixel_index:pixel_index+3] = petal_color
            
            # If this is a flower center
            if (x, y) in flowers_1:
                pixels[pixel_index:pixel_index+3] = FLOWER_CENTER
            # If this is adjacent to a flower (for petals)
            elif any((abs(x - fx) == 1 and y == fy) or (x == fx and abs(y - fy) == 1) 
                    for fx, fy in flowers_1):
                # Add some variation to the flower petals
                petal_color = add_color_noise(FLOWER_PURPLE, 5)
                pixels[pixel_index:pixel_index+3] = petal_color
    
    return pixels

if __name__ == "__main__":
    # Get tile type from command line argument
    import sys
    if len(sys.argv) != 2:
        print("Usage: python generator.py <tile_type>")
        print("Available types: grass, dirt, mixed, flowers")
        sys.exit(1)

    tile_type = sys.argv[1]
    
    # Generate requested tile
    if tile_type == "grass":
        create_tile(generate_grass_tile(), "grass_tile")
    elif tile_type == "dirt":
        create_tile(generate_dirt_tile(), "dirt_tile") 
    elif tile_type == "mixed":
        create_tile(generate_mixed_tile(), "mixed_tile")
    elif tile_type == "flowers":
       create_tile(generate_grass_with_flowers(), "flowers_tile")
    else:
        print("Invalid tile type. Available types: grass, dirt, mixed, flowers")
        sys.exit(1)
