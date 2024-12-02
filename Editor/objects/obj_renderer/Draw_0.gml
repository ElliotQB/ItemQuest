var camera_x = obj_camera.x;
var camera_y = obj_camera.y;
var camera_zoom = obj_camera.camera_zoom;
var tl_x = camera_x - (display_get_width()/camera_zoom/2);
var tl_y = camera_y - (display_get_height()/camera_zoom/2);
var br_x = tl_x + (display_get_width()/camera_zoom);
var br_y = tl_y + (display_get_height()/camera_zoom);

var step_size = 32;

draw_set_alpha(0.2);
for (var i = step_down(tl_x, step_size); i <= step_down(br_x, step_size); i+=step_size)
{
	draw_line(i, tl_y, i, br_y);
}

for (var i = step_down(tl_y, step_size); i <= step_down(br_y, step_size); i+=step_size)
{
	draw_line(tl_x, i, br_x, i);
}


draw_line(tl_x, 0, br_x, 0);
draw_line(0, tl_y, 0, br_y);


var m_x = step_down(mouse_x, step_size);
var m_y = step_down(mouse_y, step_size);

draw_rectangle(m_x, m_y, m_x+step_size, m_y+step_size, false);