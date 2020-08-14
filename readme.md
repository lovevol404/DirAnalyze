# 统计文件夹的大小
可以列出每个文件夹的大小，按照文件夹大小提供文件夹与子文件夹的树形结构，例如
```
/home/lo 46.153G
     /home/lo/JetbrainsIdesCrack.jar 29.365G
     /home/lo/.local 13.207G
          /home/lo/.local/share 13.207G
     /home/lo/.cache 2.068G
          /home/lo/.cache/netease-cloud-music 1.275G
          /home/lo/.cache/google-chrome 405.852M
          /home/lo/.cache/JetBrains 348.937M
          /home/lo/.cache/mozilla 33.308M
          /home/lo/.cache/go-build 11.156M
          /home/lo/.cache/tracker 4.509M
          /home/lo/.cache/mesa_shader_cache 2.871M
          .................
```
输入想要查看的文件夹来查看对应文件夹下的目录结构，以空格分割可输入显示深度（默认是3，深度过大会输出内容过多），如：
```
输入想要查看的文件夹：
/home/lo 2
******************************
/home/lo 46.153G
     /home/lo/JetbrainsIdesCrack.jar 29.365G
     /home/lo/.local 13.207G
     /home/lo/.cache 2.068G
     /home/lo/下载 544.837M
     /home/lo/snap 388.689M
     /home/lo/.config 194.232M
     /home/lo/Downloads 173.570M
     /home/lo/.GoLand2019.3 96.824M
     /home/lo/.IntelliJIdea2019.3 58.143M
     .............
```


