var margin_x = display_get_width() - (display_get_width()/camera_zoom);
var margin_y = display_get_height() - (display_get_height()/camera_zoom);

var camera_set_x = (x)-(display_get_width()/camera_zoom/2);
var camera_set_y = (y)-(display_get_height()/camera_zoom/2);

camera_set_view_size(camera, display_get_width()/camera_zoom, display_get_height()/camera_zoom);
camera_set_view_pos(camera, camera_set_x, camera_set_y);
camera_tar_zoom = clamp(camera_tar_zoom, 0.5, 5);
camera_zoom += (camera_tar_zoom - camera_zoom) * 0.05;

if (mouse_check_button_pressed(mb_middle) || keyboard_check_pressed(vk_control))
{
	camera_drag_offset_x = x;
	camera_drag_offset_y = y;
	
	mouse_grab_x = device_mouse_x_to_gui(0);
	mouse_grab_y = device_mouse_y_to_gui(0);
}

if (mouse_check_button(mb_middle) || keyboard_check(vk_control))
{
	mouse_drag_x = device_mouse_x_to_gui(0);
	mouse_drag_y = device_mouse_y_to_gui(0);
	
	x = camera_drag_offset_x + ((mouse_grab_x - mouse_drag_x)/camera_zoom);
	y = camera_drag_offset_y + ((mouse_grab_y - mouse_drag_y)/camera_zoom);
}

global.left_edge = x - (display_get_width()/2/camera_zoom);
global.right_edge = x + (display_get_width()/2/camera_zoom);
global.top_edge = y - (display_get_height()/2/camera_zoom);
global.bottom_edge = y + (display_get_height()/2/camera_zoom);