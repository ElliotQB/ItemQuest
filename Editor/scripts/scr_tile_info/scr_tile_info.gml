function tile_info(tile)
{
	switch (tile)
	{
		case TILES.COLLISION:
			return {
				color: c_gray,
				name: "Wall",
			};
		
		case TILES.PLAYERSTART:
			return {
				color: c_green,
				name: "Player SPos",
			};
		
		case TILES.COLLECTABLE:
			return {
				color: c_aqua,
				name: "Collectible",
			};
		
		case TILES.HAZARD:
			return {
				color: c_red,
				name: "Hazard",
			};
			
		case TILES.DOUBLEJUMPGEM:
			return {
				color: c_olive,
				name: "Double Jump Gem",
			};
			
		case TILES.TRIPLEJUMPGEM:
			return {
				color: c_navy,
				name: "Triple Jump Gem",
			};
			
		case TILES.WALLJUMPGEM:
			return {
				color: c_maroon,
				name: "Wall Jump Gem",
			};
			
			
	}
}