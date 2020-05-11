[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh
# declare -f ish_help_repos &>/dev/null || require conf.sh

require help.sh
require miss.sh

ish_miss_prepare_compile
ish_miss_prepare_install

ish_miss_prepare golang/protobuf
# ish_miss_prepare protocolbuffers/protobuf

ish_miss_prepare_volcanos
ish_miss_prepare_icebergs
ish_miss_prepare_intshell

require src/project/project.sh
require src/compile/compile.sh
require src/runtime/runtime.sh

