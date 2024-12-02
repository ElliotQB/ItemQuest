if (keyboard_check(vk_control))
{
	camera_tar_zoom -= 0.2;
}
else if (keyboard_check(vk_shift))
{
	x += 48;
}
else
{
	y += 48;
}