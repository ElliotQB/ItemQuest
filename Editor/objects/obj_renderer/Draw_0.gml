var camera_x = obj_camera.x;
var camera_y = obj_camera.y;
var camera_zoom = obj_camera.camera_zoom;
var tl_x = camera_x - (display_get_width()/camera_zoom/2);
var tl_y = camera_y - (display_get_height()/camera_zoom/2);
var br_x = tl_x + (display_get_width()/camera_zoom);
var br_y = tl_y + (display_get_height()/camera_zoom);

var step_size = 32;

var tile_data = tile_info(obj_placer.cur_tile)

// draw grid lines
draw_set_alpha(0.2);
for (var i = step_down(tl_x, step_size); i <= step_down(br_x, step_size); i+=step_size)
{
	draw_line(i, tl_y, i, br_y);
}

for (var i = step_down(tl_y, step_size); i <= step_down(br_y, step_size); i+=step_size)
{
	draw_line(tl_x, i, br_x, i);
}
draw_set_alpha(1);


// draw tiles
for (var i = step_down(global.left_edge, 32)/32; i < step_down(global.right_edge+32, 32)/32; i++)
{
	for (var j = step_down(global.top_edge, 32)/32; j < step_down(global.bottom_edge+32, 32)/32; j++)
	{
		if (i >= 0 && j >= 0 && i < 200 && j < 200)
		{
			if (obj_placer.tiles[i][j] != 0)
			{
				var _data = tile_info(obj_placer.tiles[i][j]);
				draw_set_color(_data.color);
				draw_rectangle(i*32, j*32, (i*32)+32, (j*32)+32, false);
			}
		}
	}
}
draw_set_color(c_white);


// draw axis lines
draw_line_width(tl_x, 0, br_x, 0, 4);
draw_line_width(0, tl_y, 0, br_y, 4);


// draw tile at mouse
draw_set_alpha(0.5);

var m_x = step_down(mouse_x, step_size);
var m_y = step_down(mouse_y, step_size);

draw_set_halign(fa_center);
draw_text(m_x+(step_size/2), m_y-step_size, tile_data.name);

draw_set_color(tile_data.color);

draw_rectangle(m_x, m_y, m_x+step_size, m_y+step_size, false);

draw_set_color(c_white);
draw_set_alpha(1);