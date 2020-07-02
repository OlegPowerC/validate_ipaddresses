package ipaddresses_validate

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func CheckSingleIp(IPaddr string) (err error) {
	regExString := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
	Re := regexp.MustCompile(regExString)
	iind := Re.FindStringIndex(IPaddr)
	if iind != nil && len(iind) > 0 {
		DigitVal := strings.Split(IPaddr, ".")
		for DigitIndex, DigitOne := range DigitVal {
			convval, converr := strconv.Atoi(DigitOne)
			if converr != nil {
				return errors.New("Invalid IP address")
			} else {
				if DigitIndex == 0 {
					if convval == 0 {
						return errors.New("Invalid IP address")
					}
				}
				if convval > 255 {
					return errors.New("Invalid IP address")
				}
			}
		}

		return nil
	} else {
		return errors.New("Invalid IP address")
	}
}

func MakeListIPAddresses(IPaddrs string) (IPaddrsList []string, err error) {
	var ipsplitted []string
	IPaddrs = strings.TrimSpace(IPaddrs)
	ipsplitted2 := strings.Split(IPaddrs, ",")
	for _, vars := range ipsplitted2 {
		regExString := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
		Re := regexp.MustCompile(regExString)
		iind1 := Re.FindStringIndex(IPaddrs)

		regExString = "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}-[0-9]{1,3}$"
		Re = regexp.MustCompile(regExString)
		iind2 := Re.FindStringIndex(IPaddrs)

		if iind1 == nil && iind2 == nil {
			return ipsplitted, errors.New("Invalid IP address")
		}

		if strings.Index(vars, "-") != -1 {
			//Диапазон указывается например так: 192.168.0.10-220
			vst := strings.Split(vars, "-")

			EndRange, EndRangeErr := strconv.Atoi(vst[1])
			vsi := strings.Split(vst[0], ".")
			firstipinrange_lastoctet, firstipinrange_lastoctetErr := strconv.Atoi(vsi[3])
			if firstipinrange_lastoctetErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			}

			if EndRangeErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			} else {
				if EndRange == 0 || EndRange > 255 || EndRange <= firstipinrange_lastoctet {
					return ipsplitted, errors.New("Invalid IP address")
				}
			}

			IPPrefix := vsi[0] + "." + vsi[1] + "." + vsi[2] + "."
			FirstLoctet, _ := strconv.Atoi(vsi[len(vsi)-1])
			LastLoctet, _ := strconv.Atoi(vst[1])
			for vsit := FirstLoctet; vsit <= LastLoctet; vsit++ {
				ip := IPPrefix + strconv.Itoa(vsit)
				ipsplitted = append(ipsplitted, ip)
			}
		} else {
			ipsplitted = append(ipsplitted, vars)
		}
	}
	return ipsplitted, nil
}
