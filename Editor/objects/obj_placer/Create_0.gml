#macro ARRAY_SIZE 200

tiles = [];

cur_tile = TILES.COLLISION

#macro NUM_TILES 3
enum TILES 
{
	COLLISION	= 1,
	PLAYERSTART = 2,
	COLLECTABLE = 3,
}

tile_type = 1;


for (var i = 0; i < ARRAY_SIZE; i++)
{
	array_push(tiles, array_create(ARRAY_SIZE));
}


function create_save()
{
	save = "";
	for (var i = 0; i < array_length(tiles); i++)
	{
		for (var j = 0; j < array_length(tiles[i]); j++)
		{
			if (tiles[i][j] != 0)
			{
				save += string(i) + "," + string(j) + "," + string(tiles[i][j]) + ",";
			}
		}
	}
	return save;
}