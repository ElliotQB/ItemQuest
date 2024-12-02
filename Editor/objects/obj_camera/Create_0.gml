camera = camera_create();

var vm = matrix_build_lookat(x, y, -10, x, y, 0, 0, 1, 0);
var pm = matrix_build_projection_ortho(display_get_width(), display_get_height(), 1, 10000);

camera_set_view_mat(camera, vm);
camera_set_proj_mat(camera, pm);

view_camera[0] = camera;

camera_zoom = 2;
camera_tar_zoom = 2;

mouse_grab_x = mouse_x;
mouse_grab_y = mouse_y;

mouse_drag_x = mouse_x;
mouse_drag_y = mouse_y;

camera_drag_offset_x = 0;
camera_drag_offset_y = 0;

global.left_edge = 0;
global.right_edge = 0;
global.top_edge = 0;
global.bottom_edge = 0;