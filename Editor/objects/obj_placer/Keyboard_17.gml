if (keyboard_check_pressed(ord("S")))
{
	clipboard_set_text(create_save());
}

if (keyboard_check_pressed(ord("L")))
{
	level = clipboard_get_text();
	
	split = string_split(level, ",");
	
	tiles = [];


	for (var i = 0; i < ARRAY_SIZE; i++)
	{
		array_push(tiles, array_create(ARRAY_SIZE));
	}
	
	for (var i = 0; i < array_length(split)-1; i += 3)
	{
		tiles[split[i]][split[i+1]] = split[2]
	}
}