package main
   
import (
    "fmt"
    "strconv"
    "os/exec"
    "github.com/donnie4w/go-logger/logger"
	"github.com/garyburd/redigo/redis"
)

func init(){
	logger.SetRollingFile("/data/work/src/logs", "ffmpeg.log", 10, 5, logger.MB)
}
func main() {
	c, err := redis.Dial("tcp", "115.28.241.202:7380",redis.DialPassword("6da192c7dd56a5ba917c59d2e72nneo2"))
	if err != nil{
		logger.Info(err);
	}
	defer c.Close();
	//c.Do("lpush", "video_list", "/data/work/src/ffmpeg/1.amr")
	//取列表
	values, _ := redis.Values(c.Do("LRANGE", "video_amr_list", "0", "-1"))
	for i, v := range values {
		//文件路径要是绝对路径
		from_path := string(v.([]byte));					//源路径
		to_path := "/data/work/src/video/"+strconv.Itoa(i)+".mp3";		//转码后的路径
		cmd := exec.Command("ffmpeg","-i",from_path,to_path)
		output, err := cmd.CombinedOutput()
		if err != nil {
			logger.Error(fmt.Sprint(err) + ": " + string(output));
		} else {
			logger.Info(string(output));
			c.Do("lpop","video_amr_list");			//成功弹掉一个值
			c.Do("lpush", "video_mp3_list", to_path)		//将mp3的路径记录到mp3列表里
		}
	}
}
