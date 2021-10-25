import java.io.BufferedReader
import java.io.InputStreamReader
import java.util.*

fun main(args: Array<String>) {
  val reader = BufferedReader(InputStreamReader(System.`in`))
  
  var st = StringTokenizer(reader.readLine())
  val size = st.nextToken().toInt()
  
  st = StringTokenizer(reader.readLine())
  val road = LongArray(size - 1)
  val fuel = LongArray(size)
  
  for (i in 0 until size - 1) {
    road[i] = st.nextToken().toInt().toLong()
  }
  
  st = StringTokenizer(reader.readLine())
  
  for (i in 0 until size) {
    fuel[i] = st.nextToken().toInt().toLong()
  }
  
  var min = Int.MAX_VALUE.toLong()
  var result: Long = 0
  for (i in 1 until size) {
    min = min.coerceAtMost(fuel[i - 1])
    result += road[i - 1] * min
  }
  println(result)
}