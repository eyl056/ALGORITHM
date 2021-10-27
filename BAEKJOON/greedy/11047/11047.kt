import java.io.BufferedReader
import java.io.InputStreamReader

fun `11047`() = with(BufferedReader(InputStreamReader(System.`in`))) {
  var (size: Int, target: Int) = readLine()
    .split(" ")
    .map { it.toInt() }
  val arr = mutableListOf<Int>()
  
  repeat(size) {
    arr.add(readLine().toInt())
  }
  
  if (target < 5 || size == 1) {
    print(target)
    return
  }
  var result = 0
  for (i in size - 1 downTo 0) {
    val tmp = target - arr[i]
    if (tmp == 0) {
      result++
      break
    } else if (tmp > 0) {
      val temp = target / arr[i]
      result += temp
      target -= temp * arr[i]
    } else {
      continue
    }
  }
  println(result)
}