# xmlparse 
xml config parse example

### chose xml reason  
- no mysql
配置文件，如果用mysql 每次新增字段等情况，带来复杂度提升
可以用nosql db 类似mongodb .
用xml 相对比较灵活
- no json 
是因为 json 注释不太方便，有的用_comment,也不好使. 
旧的注释数据（后边还要继续用）保留，json不支持
- TODO 
 引入配置中心，改完的xml部署后，有历史记录可查，支持回滚等操作

- target
line : map<string,string>//每行是多个key,value
file: map<string,line>// 每行id 是key
file：map<string,list<line>> // 针对 成就 、7日签到、升级 等一系列有关联的配置

支持get("file+id").getstring("field")  // 基于某个道具id，获取它的某个属性配置
支持getgroup("file-vip")=>list<line>
