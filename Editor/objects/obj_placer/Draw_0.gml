for (var i = step_down(global.left_edge, 32)/32; i < step_down(global.right_edge+32, 32)/32; i++)
{
	for (var j = step_down(global.top_edge, 32)/32; j < step_down(global.bottom_edge+32, 32)/32; j++)
	{
		if (i >= 0 && j >= 0 && i < 200 && j < 200)
		{
			if (tiles[i][j] != 0)
			{
				draw_rectangle(i*32, j*32, (i*32)+32, (j*32)+32, false);
			}
		}
	}
}