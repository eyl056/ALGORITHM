import java.io.BufferedReader
import java.io.InputStreamReader
import java.util.*

fun `11399`() {
  val reader = BufferedReader(InputStreamReader(System.`in`))
  
  var st = StringTokenizer(reader.readLine())
  val size = st.nextToken().toInt()
  
  st = StringTokenizer(reader.readLine())
  val person = mutableListOf<Int>()
  
  for (i in 0 until size) {
    person.add(st.nextToken().toInt())
  }
  
  var temp = 0
  var result = 0
  person.sort()
  for (i in 0 until size) {
    temp += person[i]
    result += temp
  }
  println(result)
}