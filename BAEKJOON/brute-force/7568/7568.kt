import java.io.BufferedReader
import java.io.InputStreamReader

fun `7568`() = with(BufferedReader(InputStreamReader(System.`in`))) {
  val size: Int = readLine().toInt()
  val arr = mutableListOf<List<Int>>()
  val rank = mutableListOf<Int>()
  
  for (i in 0 until size) {
    arr.add(
      readLine()
        .split(" ")
        .map { it.toInt() }
    )
    rank.add(1)
  }
  
  for (i in 0 until size) {
    for (j in 0 until size) {
      if (arr[i][0] < arr[j][0] && arr[i][1] < arr[j][1])
        rank[i]++
    }
  }
  
  for (i in 0 until size) {
    print("${rank[i]} ")
  }
}