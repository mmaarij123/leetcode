defmodule Solution do
  @spec remove_covered_intervals(intervals :: [[integer]]) :: integer
  def remove_covered_intervals(intervals) do
    intervals =
      Enum.sort(intervals, fn [l1, r1], [l2, r2] ->
        if l1 == l2 do
          r1 >= r2
        else
          l1 <= l2
        end
      end)

    {_max_end, count} =
      Enum.reduce(intervals, {-1, 0}, fn [_l, r], {max_end, count} ->
        if r <= max_end do
          {max_end, count}
        else
          {r, count + 1}
        end
      end)

    count
  end
end