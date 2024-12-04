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
			
			
	}
}