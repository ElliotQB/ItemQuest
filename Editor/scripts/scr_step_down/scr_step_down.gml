function step_down(val, step_size)
{
	val /= step_size;
	val = floor(val);
	return val * step_size;
}