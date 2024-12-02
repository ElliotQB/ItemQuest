var m_x = step_down(mouse_x, 32)/32;
var m_y = step_down(mouse_y, 32)/32;

if (mouse_check_button(mb_left))
{
	if (m_x >= 0 && m_y >= 0 && m_x < 200 && m_y < 200)
	{
		if (tiles[m_x][m_y] == 0)
		{
			tiles[m_x][m_y] = 1;
		}
	}
}

if (mouse_check_button(mb_right))
{
	if (m_x >= 0 && m_y >= 0 && m_x < 200 && m_y < 200)
	{
		if (tiles[m_x][m_y] != 0)
		{
			tiles[m_x][m_y] = 0;
		}
	}
}