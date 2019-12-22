package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// createAuthCookie generate a 2 hour cookie represented in uuid in redis
func createAuthCookie(username, passwd string) string {
	tuid := uuid.New().String()
	redisClient.HMSet(tuid, map[string]interface{}{
		"username": username,
		"passwd":   passwd,
	})
	redisClient.Expire(tuid, time.Hour*2)
	return tuid
}

// readByAuth returns username and passwd according to uuid
func readByAuth(tuid string) (string, string, error) {
	opt, err := redisClient.HMGet(tuid, "username", "passwd").Result()
	if err != nil {
		return "", "", err
	}
	username, ok := opt[0].(string)
	if !ok {
		return "", "", fmt.Errorf("No username ok")
	}
	passwd, ok := opt[1].(string)
	if !ok {
		return "", "", fmt.Errorf("No passwd ok")
	}
	return username, passwd, nil
}

// checkAuth checks whether auth is valid
func checkAuth(tuid string) (bool, error) {
	opt, err := redisClient.HMGet(tuid, "username", "passwd").Result()
	if err != nil {
		return false, err
	}
	_, ok := opt[0].(string)
	if !ok {
		return false, err
	}
	_, ok = opt[1].(string)
	if !ok {
		return false, err
	}
	return true, nil
}

// createComment add comment to redis with no expiration
func createComment(user, content string) bool {
	if exist, err := redisClient.Exists("commentLen").Result(); err != nil {
		fmt.Println("check exist error, ", err)
		return false
	} else if exist == 0 {
		maps := map[string]interface{}{}
		len := 0
		key := strconv.Itoa(len)
		len += 1
		keyPlus := strconv.Itoa(len)
		maps["user"] = user
		maps["content"] = content
		maps["time"] = time.Now().Format(`2006-01-02 15:04:05`)
		if _, err := redisClient.HMSet(key, maps).Result(); err != nil {
			fmt.Println("set comment error ", err)
			return false
		}
		if _, err := redisClient.Set("commentLen", keyPlus, 0).Result(); err != nil {
			fmt.Println("set commentLen error, ", err)
			return false
		}
		return true
	}
	if lenStr, err := redisClient.Get("commentLen").Result(); err != nil {
		fmt.Println("get comment len error, ", err)
		return false
	} else {
		maps := map[string]interface{}{}
		len, _ := strconv.Atoi(lenStr)
		key := strconv.Itoa(len)
		len += 1
		keyPlus := strconv.Itoa(len)
		maps["user"] = user
		maps["content"] = content
		maps["time"] = time.Now().Format(`2006-01-02 15:04:05`)
		if _, err := redisClient.HMSet(key, maps).Result(); err != nil {
			fmt.Println("set comment error ", err)
			return false
		}
		if _, err := redisClient.Set("commentLen", keyPlus, 0).Result(); err != nil {
			fmt.Println("set commentLen error, ", err)
			return false
		}
		return true
	}

}

func queryComment(curLen int) Comments {
	if exist, err := redisClient.Exists("commentLen").Result(); err != nil {
		fmt.Println("check exist error, ", err)
		return Comments{}
	} else if exist == 0 {
		redisClient.Set("commentLen", "0", 0)
		len := 0
		opt := Comments{}
		for i := curLen; i < len; i++ {
			key := strconv.Itoa(i)
			if slices, err := redisClient.HMGet(key, "user", "content", "time").Result(); err != nil {
				fmt.Println("get comment contetn errror ", err)
				return Comments{}
			} else {
				fmt.Println("slices: ", slices[0], slices[1], slices[2])
				tmp := Comment{
					Username:   slices[0].(string),
					Content:    slices[1].(string),
					CreateTime: slices[2].(string),
				}
				opt.Com = append(opt.Com, tmp)
			}
		}
		return opt
	}
	if lenStr, err := redisClient.Get("commentLen").Result(); err != nil {
		fmt.Println("get comment len error, ", err)
		return Comments{}
	} else {
		len, _ := strconv.Atoi(lenStr)
		opt := Comments{}
		for i := curLen; i < len; i++ {
			key := strconv.Itoa(i)
			if slices, err := redisClient.HMGet(key, "user", "content", "time").Result(); err != nil {
				fmt.Println("get comment contetn errror ", err)
				return Comments{}
			} else {
				fmt.Println("slices: ", slices[0], slices[1], slices[2])
				tmp := Comment{
					Username:   slices[0].(string),
					Content:    slices[1].(string),
					CreateTime: slices[2].(string),
				}
				opt.Com = append(opt.Com, tmp)
			}
		}
		return opt
	}
}

// GetCurSources returns sourceSize of public pics
func GetCurSources() int {
	if exist, err := redisClient.Exists("sourceSize").Result(); err != nil {
		fmt.Println("check exist error, ", err)
		return -1
	} else if exist == 0 {
		redisClient.Set("sourceSize", "0", 0)
		return 0
	}
	size, err := redisClient.Get("sourceSize").Result()
	if err != nil {
		fmt.Println("get image size error, ", err)
		return -1
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("parse size to sz error,", err)
		return -1
	}
	return sz
}

func addCurSources() {
	if exist, err := redisClient.Exists("sourceSize").Result(); err != nil {
		fmt.Println("check exist error, ", err)
		return
	} else if exist == 0 {
		redisClient.Set("sourceSize", "1", 0)
		return
	}
	size, err := redisClient.Get("sourceSize").Result()
	if err != nil {
		fmt.Println("get source size error, ", err)
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("parse size to sz error,", err)
	}
	sz += 1
	szstr := strconv.Itoa(sz)
	redisClient.Set("sourceSize", szstr, 0)
}

func newSource(sourceName, realName string) {
	dbkey := "source@" + sourceName
	maps := make(map[string]interface{})
	maps["realName"] = realName
	maps["postTime"] = time.Now().Format(`2006-01-02 15:04:05`)
	maps["liked"] = "0"
	redisClient.HMSet(dbkey, maps)
}

func readSource(sourceID string) PublicPhotoInfo {
	dbKey := "source@" + sourceID
	maps, err := redisClient.HMGet(dbKey, "realName", "postTime", "liked").Result()
	if err != nil {
		fmt.Println(err)
	}
	liked, _ := strconv.Atoi(maps[2].(string))
	return PublicPhotoInfo{
		IdName:   sourceID,
		RealName: maps[0].(string),
		PostTime: maps[1].(string),
		Liked:    liked,
	}
}
