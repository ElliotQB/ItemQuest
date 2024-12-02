tiles = [];


var array_size = 200;
for (var i = 0; i < array_size; i++)
{
	array_push(tiles, array_create(array_size));
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